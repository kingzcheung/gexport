create table `a` (
    id int not null primary key comment 'id',
    name varbinary(500) comment 'test name',
    unique (name)
) comment 'test table'