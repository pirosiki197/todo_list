CREATE TABLE `todos` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `task` VARCHAR(255) NOT NULL,
    `status` VARCHAR(32) NOT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
