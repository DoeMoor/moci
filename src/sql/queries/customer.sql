-- name: GetAllCustomers :many
select id, customers.name from customers;

-- name: GetAllCustomersWithProjects :many
select c.id, c.name, pj.name
from customers c
left join customer_projects_junction cpj on c.id = cpj.customers_id
left join projects pj on cpj.projects_id = pj.id;

-- name: GetCustomerWithProjects :many
select c.id, c.name, pj.name
from customers c
left join customer_projects_junction cpj on c.id = cpj.customers_id
left join projects pj on cpj.projects_id = pj.id
where c.id = $1;

