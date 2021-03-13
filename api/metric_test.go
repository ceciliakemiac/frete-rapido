package api

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/ceciliakemiac/frete-rapido/model"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/gorm"
)

func createTables(pg *gorm.DB) {
	pg.AutoMigrate(&model.Freight{})
}

func dropTables(pg *gorm.DB) {
	pg.Migrator().DropTable(&model.Freight{})
}

func fillTables(pg *gorm.DB) {
	freight1 := model.Freight{
		ID:           1,
		Nome:         "FR EXPRESS (TESTE)",
		Servico:      "Normal",
		PrazoEntrega: 5,
		PrecoFrete:   72.5,
		QuoteID:      1,
	}
	freight2 := model.Freight{
		ID:           2,
		Nome:         "RAPIDÃO FR (TESTE)",
		Servico:      "Normal",
		PrazoEntrega: 5,
		PrecoFrete:   74.25,
		QuoteID:      1,
	}
	freight3 := model.Freight{
		ID:           3,
		Nome:         "FR EXPRESS (TESTE)",
		Servico:      "Express",
		PrazoEntrega: 4,
		PrecoFrete:   87.91,
		QuoteID:      2,
	}
	freight4 := model.Freight{
		ID:           4,
		Nome:         "EXPRESSO FR (TESTE)",
		Servico:      "Normal",
		PrazoEntrega: 5,
		PrecoFrete:   101.17,
		QuoteID:      2,
	}

	freights := []model.Freight{}
	freights = append(freights, freight1)
	freights = append(freights, freight2)
	freights = append(freights, freight3)
	freights = append(freights, freight4)

	pg.Create(freights)
}

func TestGetMetricsNoLastQuotes(t *testing.T) {
	server, mock := InitServer()

	dropTables(server.db.PG)
	createTables(server.db.PG)
	fillTables(server.db.PG)

	mock.ExpectQuery(regexp.QuoteMeta(
		`select count(nome) as total_ocorrencias, nome as transportadora, 
		sum(preco_frete) as total_precos, 
		cast(sum(preco_frete) as float) / cast(count(nome) as float) as media_precos
		from freights
		group by transportadora
		order by total_ocorrencias desc`,
	)).
		WillReturnRows(sqlmock.NewRows([]string{"total_ocorrencias", "transportadora", "total_precos", "media_precos"}).AddRow(2, "FR EXPRESS (TESTE)", 160.41, 80.205).AddRow(1, "RAPIDÃO FR (TESTE)", 74.25, 74.25).AddRow(1, "EXPRESSO FR (TESTE)", 101.17, 101.17))

	mock.ExpectQuery(regexp.QuoteMeta(
		`select nome, servico, prazo_entrega, preco_frete as valor
		from freights 
		where preco_frete = (select min(preco_frete) from freights)`,
	)).
		WillReturnRows(sqlmock.NewRows([]string{"nome", "servico", "prazo_entrega", "preco_frete"}).AddRow("FR EXPRESS (TESTE)", "Normal", 5, 72.5))

	mock.ExpectQuery(regexp.QuoteMeta(
		`select nome, servico, prazo_entrega, preco_frete as valor
			from freights 
			where preco_frete = (select max(preco_frete) from freights)`,
	)).
		WillReturnRows(sqlmock.NewRows([]string{"nome", "servico", "prazo_entrega", "preco_frete"}).AddRow("EXPRESSO FR (TESTE)", "Normal", 5, 101.17))

	w := httptest.NewRecorder()
	server.router.ServeHTTP(w, httptest.NewRequest("GET", "/api/metrics", nil))

	if w.Code != http.StatusOK {
		t.Error("Did not get expected HTTP status code, got", w.Code)
	}

	if w.Body.String() != `{
		"fretes":[
			{
				"transportadora":"FR EXPRESS (TESTE)",
				"total_ocorrencias":2,
				"total_precos":160.41,
				"media_precos":80.205
			},
			{
				"transportadora":"RAPIDÃO FR (TESTE)",
				"total_ocorrencias":1,
				"total_precos":74.25,
				"media_precos":74.25
			},
			{
				"transportadora":"EXPRESSO FR (TESTE)",
				"total_ocorrencias":1,
				"total_precos":101.17,
				"media_precos":101.17
			}
			],
			"frete_mais_caro":
			{
				"nome":"EXPRESSO FR (TESTE)",
				"servico":"Normal",
				"prazo_entrega":5,
				"valor":101.17
			},
			"frete_mais_barato":
			{
				"nome":"FR EXPRESS (TESTE)",
				"servico":5,
				"prazo_entrega":"Normal",
				"valor":72.5
			}
		}` {
		t.Error("Did not get right response, got ", w.Body.String())
	}
}

func TestMetricsLastQuotes(t *testing.T) {
	server, mock := InitServer()

	dropTables(server.db.PG)
	createTables(server.db.PG)
	fillTables(server.db.PG)

	mock.ExpectQuery(regexp.QuoteMeta(
		`select count(nome) as total_ocorrencias, nome as transportadora, 
		sum(preco_frete) as total_precos, 
		cast(sum(preco_frete) as float) / cast(count(nome) as float) as media_precos
		from freights 
		where quote_id in (
			select distinct quote_id 
			from freights 
			order by quote_id desc 
			limit 1
		)
		group by transportadora
		order by total_ocorrencias desc`,
	)).
		WillReturnRows(sqlmock.NewRows([]string{"total_ocorrencias", "transportadora", "total_precos", "media_precos"}).AddRow(2, "FR EXPRESS (TESTE)", 87.91, 87.91).AddRow(1, "EXPRESSO FR (TESTE)", 101.17, 101.17))

	mock.ExpectQuery(regexp.QuoteMeta(
		`select f.nome, f.servico, f.prazo_entrega, f.preco_frete as valor
			from (
				select *
				from freights 
				where quote_id in (
					select distinct quote_id
					from freights 
					order by quote_id desc 
					limit 1
				)
			) as f
			where f.preco_frete = (select min(e.preco_frete) from (
				select *
				from freights 
				where quote_id in (
					select distinct quote_id
					from freights 
					order by quote_id desc 
					limit 1
				)
			) as e)
			`,
	)).
		WillReturnRows(sqlmock.NewRows([]string{"nome", "servico", "prazo_entrega", "preco_frete"}).AddRow("FR EXPRESS (TESTE)", "Express", 4, 87.91))

	mock.ExpectQuery(regexp.QuoteMeta(
		`select f.nome, f.servico, f.prazo_entrega, f.preco_frete as valor
			from (
				select *
				from freights 
				where quote_id in (
					select distinct quote_id
					from freights 
					order by quote_id desc 
					limit 1
				)
			) as f
			where f.preco_frete = (select max(e.preco_frete) from (
				select *
				from freights 
				where quote_id in (
					select distinct quote_id
					from freights 
					order by quote_id desc 
					limit 1
				)
			) as e)
	`,
	)).
		WillReturnRows(sqlmock.NewRows([]string{"nome", "servico", "prazo_entrega", "preco_frete"}).AddRow("EXPRESSO FR (TESTE)", "Normal", 5, 101.17))

	w := httptest.NewRecorder()
	server.router.ServeHTTP(w, httptest.NewRequest("GET", "/api/metrics?last_quotes=1", nil))

	if w.Code != http.StatusOK {
		t.Error("Did not get expected HTTP status code, got", w.Code)
	}

	if w.Body.String() != `{
		"fretes": [
			{
				"transportadora": "EXPRESSO FR (TESTE)",
				"total_ocorrencias": 1,
				"total_precos": 101.17,
				"media_precos": 101.17
			},
			{
				"transportadora": "FR EXPRESS (TESTE)",
				"total_ocorrencias": 1,
				"total_precos": 87.91,
				"media_precos": 87.91
			}
		],
		"frete_mais_caro": {
			"nome": "EXPRESSO FR (TESTE)",
			"servico": "Normal",
			"prazo_entrega": 5,
			"valor": 101.17
		},
		"frete_mais_barato": {
			"nome": "FR EXPRESS (TESTE)",
			"servico": "Express",
			"prazo_entrega": 4,
			"valor": 87.91
		}
	}` {
		t.Error("Did not get right response, got ", w.Body.String())
	}
}
