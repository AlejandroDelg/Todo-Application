create table todos(
    id serial not null,
    name varchar(255),
    IsDone  bool,
    primary key(id)
)

insert into todos(todo, IsDone) values ("First todo", false)