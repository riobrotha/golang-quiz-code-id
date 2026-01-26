
select shipper_id, company_name, product_id, product_name, max(total_qty_ordered) from (
  select
  s.shipper_id,
  s.company_name,
  p.product_id,
  p.product_name,
  count(od.order_id) as total_qty_ordered
  from oe.shippers s
  left join oe.orders o on s.shipper_id = o.ship_via
  left join oe.order_details od on o.order_id = od.order_id
  left join oe.products p on od.product_id = p.product_id
  group by s.shipper_id, s.company_name, p.product_id, p.product_name
) as t
group by shipper_id, company_name, product_id, product_name
UNION ALL
select shipper_id, company_name, product_id, product_name, min(total_qty_ordered) from (
  select
  s.shipper_id,
  s.company_name,
  p.product_id,
  p.product_name,
  count(od.order_id) as total_qty_ordered
  from oe.shippers s
  left join oe.orders o on s.shipper_id = o.ship_via
  left join oe.order_details od on o.order_id = od.order_id
  left join oe.products p on od.product_id = p.product_id
  group by s.shipper_id, s.company_name, p.product_id, p.product_name
) as t
group by shipper_id, company_name, product_id, product_name
order by shipper_id

