BEGIN;


CREATE TABLE users (
    id SERIAL NOT NULL,
    email varchar(100) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    last_login timestamp,
    PRIMARY KEY (id)
);

CREATE TABLE user_profiles(
    user_id int NOT NULL,
    first_name varchar(100) NOT NULL,
    last_name varchar(100) NOT NULL,
    phone_number varchar(100),
    CONSTRAINT FK_UserProfile FOREIGN KEY (user_id)
    REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE role_types(
    id SERIAL NOT NULL,
    name varchar(30) NOT NULL,
    level int NOT NULL default 1,
    PRIMARY KEY (id)
);

CREATE TABLE user_roles(
    user_id int NOT NULL,
    role_id int NOT NULL,
    CONSTRAINT FK_UserRole FOREIGN KEY (user_id)
    REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT FK_RoleType FOREIGN KEY (role_id)
    REFERENCES role_types(id) ON DELETE CASCADE
);

CREATE TABLE projects (
    id SERIAL NOT NULL,
    name varchar(100) NOT NULL UNIQUE,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE flags (
    id SERIAL NOT NULL,
    project_id int NOT NULL,
    name varchar(100) NOT NULL UNIQUE,
    active boolean DEFAULT false,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT FK_ProjectFlag FOREIGN KEY (project_id)
    REFERENCES projects(id) ON DELETE CASCADE
);

CREATE TABLE flag_contexts (
  id SERIAL NOT NULL,
  flag_id int NOT NULL,
  name varchar(100),
  condition varchar(10),
  value varchar(100),
  update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT FK_FlagContext FOREIGN KEY (flag_id)
  REFERENCES flags(id) ON DELETE CASCADE
);

INSERT INTO role_types (name, level)
VALUES ('Low Staff', 1);

INSERT INTO role_types (name, level)
VALUES ('Middle Staff', 2);

INSERT INTO role_types (name, level)
VALUES ('High Staff', 3);

INSERT INTO role_types (name, level)
VALUES ('Administrator', 10);

COMMIT;