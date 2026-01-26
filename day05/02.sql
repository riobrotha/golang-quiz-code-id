
select 
s.supplier_id,
s.company_name,
count(p.product_id) as total_product
from oe.suppliers s
left join oe.products p on s.supplier_id = p.supplier_id
group by s.supplier_id, s.company_name
order by total_product desc
