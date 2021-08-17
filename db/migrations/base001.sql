create table txnentry(
id serial,
entity_id varchar(64) not null,
created_at timestamp not null,
value NUMERIC(19,3)
);

create index idx_created_at on txnentry(created_at);

create index idx_entity_id on txnentry(entity_id);