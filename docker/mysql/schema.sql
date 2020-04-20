CREATE DATABASE secretform;
CREATE TABLE secretform.forms (
    ID VARCHAR(40) PRIMARY KEY,
    Content TEXT,
    PublicKey TEXT,
    Open BOOLEAN NOT NULL
);

CREATE TABLE secretform.answers (
    ID VARCHAR(40) PRIMARY KEY,
    Data TEXT,
    Form VARCHAR(40) NOT NULL,
    CONSTRAINT fk_form FOREIGN KEY secretform.answers(Form) REFERENCES secretform.forms(ID);
);

CREATE USER dev@'%' IDENTIFIED BY 'dev';
GRANT ALL PRIVILEGES ON secretform.* TO dev@'%';
