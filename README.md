## Enigma Mart

### Database
```sql
CREATE database db_enigma_mart;

CREATE TABLE m_product (
    id varchar(60) primary key,
    name varchar(100) not null,
    price int,
    stock int,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp
);

CREATE TABLE t_product_sell (
    id varchar(60) primary key,
    date date,
    factur varchar(60),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp
);

CREATE TABLE t_product_sell_detail (
    id varchar(60) primary key,
    product_sell_id varchar(60),
    product_id varchar(60),
    qty int,
    price int,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp,
    foreign key(product_sell_id) references t_product_sell(id),
    foreign key(product_id) references m_product(id)
);
```

### Run
```
DB_DRIVER=postgres DB_HOST=localhost DB_PORT=5432 DB_USER=postgres DB_PASSWORD=P@ssw0rd DB_NAME=db_enigma_mart go run  .
```