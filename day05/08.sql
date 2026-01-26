select c.category_id, c.category_name, count(od.order_id) as total_qty_ordered
from oe.categories c
left join oe.products p on c.category_id = p.category_id
left join oe.order_details od on p.product_id = od.product_id
group by c.category_id, c.category_name
order by total_qty_ordered desc
