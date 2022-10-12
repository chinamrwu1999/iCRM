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
   Name varchar(30) not null,
   Role varchar(8) ,
   Password varchar(100) 
);

;
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
    Description varchar(300),
    createTime datetime default now()
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

create table MarketNames(    # 市场大区名：东大区、南大区、北大区等
    AreaId SMALLINT not null primary key,
    Name varchar(30) not null
);

delete from MarketNames;
INSERT INTO MarketNames VALUES(0,'东大区'),(1,'南大区'),(2,'北大区'),(3,'国际');

create table MarketProvinces ( # 各个市场大区包含的省市
   AreaId int not null,
   Code varchar(10) not null,
   primary key(AreaId,Code)
);

INSERT INTO MarketProvinces VALUES
(0,'110000'),(0,'310000'),(0,'320000'),(0,'320100'),(0,'320800'),
(0,'321000'),(0,'321100'),(0,'321200'),(0,'321300'),(0,'330000'),
(0,'340000'),(0,'370000'),(0,'410000'),(0,'420000'),(2,'120000'),
(2,'130000'),(2,'140000'),(2,'150000'),(2,'210000'),(2,'220000'),
(2,'230000'),(2,'610000'),(2,'620000'),(2,'630000'),(2,'640000'),
(2,'650000'),(1,'350000'),(1,'360000'),(1,'430000'),(1,'440100'),
(1,'440300'),(1,'450000'),(1,'460000'),(1,'500000'),(1,'510000'),
(1,'520000'),(1,'530000'),(1,'540000');




create table PositionNames(
   Code varchar(5) not null primary key,
   Name varchar(20) not null
);
delete from PositionNames;
INSERT INTO PositionNames values('SS','大区总监'),('PM','省区经理'),('MM','招商推广经理'),('SM','销售主管');

create table MarketPersons(
   EmployeeId varchar(10) not null, # 员工号
   Code varchar(10) not null, # 行政区号
   Status TINYINT(1), # 状态：0离职,0 在职
   startDate dateTime default now(), # 开始日期
   primary key(EmployeeId,Code)
);

create table Hospitals(
    ID int auto_increment not null primary key,
    Name varchar(60) not null,
    Code varchar(10) not  null,
    Grade varchar(10),
    HType varchar(10)
);
