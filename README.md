# About this project.

I was looking for an easy way to store my notes and tech tips, so I thought, why not build a custom bot on telegram?

This project shows step by step how I built my bot with the following technologies:
- Raspberry pi;
- Docker;
- MariaDB;
- Golang;

![image](https://user-images.githubusercontent.com/12565936/169705514-09bda6f1-1411-4b28-b1ea-9fb85fc2ecf8.png)


# How it works?
- Annotations are divided into categories and subcategories. This is useful for navigating between topics.
- In the database, we have a table with the categories (subject), we also have a table with the subcategories (tip) and with the notes about subtopic.
- When selecting a subcategory, the bot will show the annotation and will also show an example image if it exists.

Example: In the `git` category, we have the subcategory that brings the note on how to create an ssh key

![](https://github.com/dihr/notes/blob/main/aux_file/bot_example_with_image.gif)

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
    `flag_img` bool not null,
    `img_file_name` varchar(200),
    primary key (`id`),
    foreign key (category_id) references category(id)
)
```

# Running app image.
Once the project has already been cloned to the raspbarry py local folder, run the commands:
- Creating the image:
  -`docker build -t <image_name> . `
  ![image](https://user-images.githubusercontent.com/12565936/169706162-cebc6817-03ac-41f2-a757-ff8d60484d31.png)

- Running the image:
  -`docker run -d -v /var/lib/imgs:/images -e TELEGRAM_BOT_TOKEN="<BOT_TOKEN>" -e APP_DB_HOST=192.168.1.2 -e APP_DB_PASSWORD=<DB_PASSWORD> -e APP_DB_PORT=3306 -e APP_DB_USER=root telegram_bot`
  ![image](https://user-images.githubusercontent.com/12565936/169713198-688c6486-cf1d-47a6-a06e-d428a4448d32.png)
  
- Loading images:
  - If you want to include images, just save the files in a volume and named /images. The files must contain the same name as the subcategory.


# Adding notes.

To add new annotations, simply insert them into the category and subcategory tables.
