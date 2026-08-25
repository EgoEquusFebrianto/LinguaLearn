create table users (
	id BIGINT unsigned auto_increment primary key,
    full_name varchar(255) not null,
    email varchar(255) not null,
    password_hash varchar(255) not null,
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp
		on update current_timestamp
);