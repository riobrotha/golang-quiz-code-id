select 
p.product_id,
p.product_name,
count(od.product_id) as total_qty
from oe.products p
left join oe.order_details od on p.product_id = od.product_id
group by p.product_id, p.product_name
order by total_qty desc