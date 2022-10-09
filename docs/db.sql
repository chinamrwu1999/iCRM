create database iCRM;
use iCRM;

create table businessAreas(
     ID TINYINT not null primary key,
     Name varchar(20) not null
);

create table businessAreaProvinces(
    AreaId TINYINT not null,
    AreaCode varchar(10) not null
);

create table Departments(
   DeptId TINYINT not null,
   Name varchar(30)
);
create table Employees(
   ID char(8) not null primary key,
   DeptID TINYINT not null,
   Name varchar(30) not null,
   Gender char(1),
   Role varchar(4) not null,
   Password varchar(20) not null
);

create table Customers(
    ID int not null auto_increment primary key,
    FullName varchar(60) not  null,
    ShortName varchar(30),
    CType char(20),
    Scale char(20),
    Status char(20),
    Level char(20),
    GetWay char(20),
    Nation char(20) default 'cn',
    Province char(6),
    City Char(6),
    Address varchar(100),
    Description varchar(300)
);

create table CustomerContacts(
    CustomerId int not null,
    Name varchar(30) not null,
    Email varchar(50),
    Phone varchar(20),
    updateTime datetime not null default now(),
    primary key(CustomerId,Phone)
);
Create table CustomerBanks(
    CustomerId int not null primary key,
    BankName varchar(200),
    BankAccount varchar(60),
    AccountName varchar(200),
    TaxID varchar(60)
);

create table Codes(
    Label varchar(100) not null,
    Code varchar(100) not null,
    CodeType varchar(30) not null,
    displayOrder SMALLINT ,
    Remark varchar(30),
    primary key(Code,CodeType)
);

create table nations(
    code varchar(20) not null primary key,
    Name varchar(50) not null
);
create table citys(
    Code varchar(10) not null primary key,
    Name varchar(50) not null,
    ParentId varchar(20)
);
