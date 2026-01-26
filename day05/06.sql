with cte_orders as (
  select
  order_id,
  customer_id,
  order_date,
  required_date,
  shipped_date,
  shipped_date - order_date as delivery_time
  from oe.orders
  order by order_id
)

select * from cte_orders where delivery_time > 7;