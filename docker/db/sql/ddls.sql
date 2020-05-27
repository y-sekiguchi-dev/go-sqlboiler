DROP TABLE IF EXISTS children;
DROP TABLE IF EXISTS persons;
DROP TABLE IF EXISTS users;

create table IF not exists users
(
    user_id         BIGINT AUTO_INCREMENT,
    first_name      VARCHAR(50) NOT NULL,
    last_name       VARCHAR(50) NOT NULL,
    PRIMARY KEY (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

create table IF not exists persons
(
    person_id       BIGINT AUTO_INCREMENT,
    first_name      VARCHAR(50) NOT NULL,
    last_name       VARCHAR(50) NOT NULL,
    birthday        DATE NOT NULL,
    personality     TINYINT,
    has_partner     TINYINT(1) NOT NULL DEFAULT 0,
    version         SMALLINT NOT NULL DEFAULT 0,
    deleted         TINYINT(1) NOT NULL DEFAULT 0,
    created_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_user_id BIGINT NOT NULL,
    updated_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_user_id BIGINT NOT NULL,
    PRIMARY KEY (person_id),
    CONSTRAINT fk_persons_created_user_id FOREIGN KEY (created_user_id) REFERENCES users (user_id),
    CONSTRAINT fk_persons_updated_user_id FOREIGN KEY (updated_user_id) REFERENCES users (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

create table IF not exists children
(
    person_id       BIGINT NOT NULL,
    sub_no          TINYINT NOT NULL,
    first_name      VARCHAR(50) NOT NULL,
    last_name       VARCHAR(50) NOT NULL,
    birthday        DATE NOT NULL,
    PRIMARY KEY (person_id, sub_no),
    CONSTRAINT fk_children_person_id FOREIGN KEY (person_id) REFERENCES persons (person_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
