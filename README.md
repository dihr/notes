# About this project.

I was looking for an easy way to store my notes and tech tips, so I thought, why not build a custom bot on telegram?

This project shows step by step how I built my bot with the following technologies:
- Raspberry pi;
- Docker;
- MariaDB;
- Golang;

# How it works?
- Annotations are divided into categories and subcategories. This is useful for navigating between topics.
- In the database, we have a table with the categories (subject), we also have a table with the subcategories (tip) and with the notes about subtopic.

Example: In the `git` category, we have the subcategory that brings the note on how to create an ssh key

# Configuring the raspberry PI.

- Install any linux distribution for ARM architecture.I ended up using the kali distribution which can be obtained from the [link](https://www.kali.org/get-kali/#kali-arm).
- Install docker using commands
  - `sudo apt update`
  - `sudo apt install -y docker.io`

# Running MariDB image.

- To run the MariaDB image use the command below changing the username and password variables.
  - `docker run -it --name mysql -p 3306:3306 -v /var/lib/mysql:/var/lib/mysql -e MYSQL_DATABASE=wordpressdb -e MYSQL_USER=wordpressuser -e MYSQL_PASSWORD=hguyFt6S95dgfR4ryb -e MYSQL_ROOT_PASSWORD=hguyFtgfR4r9R4r76 yobasystems/alpine-mariadb`

# Creating tables.
- Category:
```
create table category(
    `id` int auto_increment not null,
    `name` varchar(200) not null,
    primary key (`id`)
)
```
- Subcategory:
```
create table sub_category(
    `id` int auto_increment not null,
    `category_id` int not null,
    `name` varchar(200) not null,
    `text` text not null,
    primary key (`id`),
    foreign key (category_id) references category(id)
)
```

# Adding notes.

To add new annotations, simply insert them into the category and subcategory tables.