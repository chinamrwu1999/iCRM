create database iCRM;
use iCRM;

create table department(
   DeptId TINYINT not null,
   Name varchar(30)
);
create table employee(
   ID char(8) not null primary key,
   Name varchar(30) not null,
   Role varchar(8) ,
   Password varchar(100) 
);



create table code(
    Label varchar(100) not null,
    Code varchar(100) not null,
    CodeType varchar(30) not null,
    DisplayOrder SMALLINT ,
    Remark varchar(30),
    primary key(Code,CodeType)
);

create table nation(
    code varchar(20) not null primary key,
    Name varchar(50) not null
);
create table city(
    Code varchar(10) not null primary key,
    Name varchar(50) not null,
    ParentId varchar(20)
);

drop table marketArea;
create table marketArea(    # 市场大区名：东大区、南大区、北大区等
    AreaId varchar(2) not null primary key,
    Name varchar(30) not null
);
INSERT INTO marketArea VALUES('E','东大区'),('S','南大区'),('N','北大区');

drop table marketProvince;
create table marketProvince ( # 各个市场大区包含的省市
   AreaId char not null,
   Code varchar(10) not null,
   primary key(AreaId,Code)
);

delete from marketProvince;
INSERT INTO marketProvince VALUES
('E','110000'),('E','310000'),('E','320000'),('E','320100'),('E','320800'),('E','321000'),('E','321100'),('E','321200'),('E','321300'),('E','330000'),
('E','340000'),('E','370000'),('E','410000'),('E','420000'),
('N','120000'),('N','130000'),('N','140000'),('N','150000'),('N','210000'),('N','220000'),('N','230000'),('N','610000'),('N','620000'),
('N','630000'),('N','640000'),('N','650000'),('S','350000'),('S','360000'),('S','430000'),('S','440100'),
('S','440300'),('S','450000'),('S','460000'),('S','500000'),('S','510000'),('S','520000'),('S','530000'),('S','540000');




create table positionName(
   Code varchar(5) not null primary key,
   Name varchar(20) not null
);
delete from PositionName;
INSERT INTO PositionName values('SS','大区总监'),('PM','省区经理'),('MM','招商推广经理'),('SM','销售主管');

create table marketPerson(
   EmployeeId varchar(10) not null, # 员工号
   Code varchar(10) not null, # 行政区号
   Status TINYINT(1), # 状态：0离职,0 在职
   startDate dateTime default now(), # 开始日期
   
   primary key(EmployeeId,Code)
);

create table hospital(
    ID int auto_increment not null primary key,
    Name varchar(100) not null,
    ShortName varchar(20),
    City varchar(10) not  null,
    
    Grade varchar(10),
    HType varchar(10)
);


create table customer(
    ID int not null auto_increment primary key,
    Name varchar(60) not  null,
    ShortName varchar(30),
    City varchar(10),

    CType char(20),
    Scale char(20),
    Status char(20),
    Level char(20),
    GetWay char(20),
    Address varchar(100),
    Description varchar(300),
    CreateTime datetime default now()
);
#代理商
create table proxy(
    ID int not null auto_increment primary key,
    Name varchar(60) not  null,
    ShortName varchar(30),
     City varchar(10),
    Status char(20),
    Address varchar(100),
    CreateTime datetime default now()
);

 create table proxyCustomer(
    proxyId int not null,
    customerId int not null
    status flag varchar(4) not null default 'A' , # A:  
    createTime dateTime now(),
    primary key(proxyId,customerId)
  );

 create table proxyHospital(
    proxyId int not null,
    hospitalID int not null,
    status flag varchar(4) not null default 'A' , # A:  
    createTime dateTime now(),
    primary key(proxyId,hospitalId)
  );

create table customerContact(
    CustomerId int not null,
    Name varchar(30) not null,
    Email varchar(50),
    Phone varchar(20),
    UpdateTime datetime not null default now(),
    primary key(CustomerId,Phone)
);
Create table customerBank(
    CustomerId int not null primary key,
    BankName varchar(200),
    BankAccount varchar(60),
    AccountName varchar(200),
    TaxID varchar(60)
);

###################################
drop table businessLog;
create table businessLog(
    ID BIGINT not null auto_increment primary key,
    EmployeeId varchar(10) not null,
    hospitalId int,
    proxyId int,
    customerId int,
    content varchar(500) not null,
    stage SMALLINT,
    workingDate date default(CURRENT_DATE),
    createTime time default(time(now())),
    unique(EmployeeId,workingDate,hospitalId)
);





create table estimation( 
    customerId int not null,
    productId varchar(20) not null,
    saleYear int not null,
    saleMonth TINYINT,
    price decimal(5,2),
    amount int not null,
    submiter varchar(10),
    createTime dateTime default now(),
    primary key(customerId,productId,saleYear,saleMonth)
 );



drop table product;
create table product( 
    ID int not null auto_increment primary key,
    productId varchar(20) not null unique ,
    Name varchar(30) not null,
    basePrice decimal(7,2)
);
INSERT INTO product (productId,Name,basePrice) values
('ACK','艾长康试剂',180.0),('ACKLDT','艾长康收样检测',360.0),
('AXA','艾消安试剂',580.0),('AXALDT','艾消安收样检测',880.0);
