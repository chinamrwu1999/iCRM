create table iCRM;
use iCRM;
create table areas(
    Code varchar(10) not null primary key,
    Name varchar(80) not null,
    ParentCode varchar(10) not null,
);

create table businessAreas(
     ID TINYINT not null primary key,
     Name varchar(20) not null
);

create table businessAreaProvinces(
    AreaId TINYINT not null,
    AreaCode varchar(10) not null,
);

create table Departments(
   DeptId TINYINT not null,
   Name varchar(30)
);
create table Employees(
   ID char(8) not null primary key;
   DeptID TINYINT not null,
   Name varchar(30) not null,
   Gender char(1),
   Role varchar(4) not null,
   Password varchar(20) not null,
);


