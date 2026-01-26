select 
c.category_id,
c.category_name,
count(p.product_id) as total_product
from oe.categories c
left join oe.products p on c.category_id = p.category_id
group by c.category_id, c.category_name
order by c.category_id

