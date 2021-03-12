package database

const (
	queryGetMetrics = `
		select count(nome) as total_ocorrencias, nome as transportadora, 
		sum(preco_frete) as total_precos, 
		cast(sum(preco_frete) as float) / cast(count(nome) as float) as media_precos
		from freights
		group by transportadora
		order by total_ocorrencias desc
	`

	queryGetMetricsLastQuotes = `
		select count(nome) as total_ocorrencias, nome as transportadora, 
		sum(preco_frete) as total_precos, 
		cast(sum(preco_frete) as float) / cast(count(nome) as float) as media_precos
		from freights 
		where quote_id in (
			select distinct quote_id 
			from freights 
			order by quote_id desc 
			limit %d
		)
		group by transportadora
		order by total_ocorrencias desc
	`

	queryPrice = `
		select nome, servico, prazo_entrega, preco_frete as valor
		from freights 
		where preco_frete = (select %s(preco_frete) from freights)
	`

	queryPriceLastQuotes = `
		select f.nome, f.servico, f.prazo_entrega, f.preco_frete as valor
		from (
			select *
			from freights 
			where quote_id in (
				select distinct quote_id
				from freights 
				order by quote_id desc 
				limit %d
			)
		) as f
		where f.preco_frete = (select %s(e.preco_frete) from (
			select *
			from freights 
			where quote_id in (
				select distinct quote_id
				from freights 
				order by quote_id desc 
				limit %d
			)
		) as e)
	`
)
