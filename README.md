# mytheresa

### tests are not done - not enough time

## Starting service

`docker compose up`

After that you can call endpoint: `http://localhost:8080/api/v1/products?category=boots&priceLessThan=90000`

## Decisions

It's a little bit complicated DDD style. It gives some advantages in big projects but may be hard to read sometimes. I have tried to implement everything to be adjustible. <br>

I choose Chi but I also used Gin - don't see much difference if it's not really highload.<br>

PostgreSQL for database. Again, it can be MySQL. Or even NoSQL. But I assume service will grow and will have additional entities like discounts.<br>
Discount entity are not fully implemented. It should be somewhere in database to be configurable. And we can have a worker that will keep discounts in memory.<br>
SKU can be used as unique primary key - but I am not sure if it's unique right now and will be in future (for example 2 products with 1 SKU but different currency).<br>
Index is by category and price.

I have used "migrate" library for simplicity. Usually I use liquibase for migrations, but for MySQL, had no time to switch to PostgreSQL.

I used sqlx for DB managment. It also a good idea to use something like Masterminds/squirrel for SQL quieres managment (GetProducts method is a good example. squirell will allow not to use ifs for different number of arguments.)

TODO:<br>
Tests<br>
Metrics<br>
Masterminds/squirrel<br>
Liquibase<br>
Move discount to be separate entitity.<br>
Traces<br>
Maybe: auth<br>
