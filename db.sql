-- For database My-Pokemon Page --
CREATE TABLE `pokemon`.`my_pokemon` (
    id int NOT NULL AUTO_INCREMENT,
    pokemon_name varchar(50) NOT NULL,
    nick_name varchar(50),
    image varchar(500),
    update_count int DEFAULT 0,
    created_on timestamp DEFAULT CURRENT_TIMESTAMP
);
