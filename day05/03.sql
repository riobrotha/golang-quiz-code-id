
select 
s.supplier_id,
s.company_name,
count(p.product_id) as total_product,
to_char(AVG(p.unit_price), 'FM99999.00') as avg_unit_price
from oe.suppliers s
left join oe.products p on s.supplier_id = p.supplier_id
group by s.supplier_id, s.company_name
order by total_product desc


