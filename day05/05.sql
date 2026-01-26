select c.customer_id, c.company_name, count(o.order_id)
from oe.customers c
left join oe.orders o on c.customer_id = o.customer_id
group by c.customer_id, c.company_name
order by c.customer_id