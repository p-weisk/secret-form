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

DELIMITER //

CREATE TRIGGER check_form_open
BEFORE INSERT
   ON secretform.answers FOR EACH ROW

BEGIN

   DECLARE isOpen BOOLEAN;
   SELECT Open From secretform.forms INTO isOpen WHERE ID = NEW.Form;

    IF isOpen = false THEN
        SIGNAL sqlstate '45000' SET message_text = 'FORM_IS_CLOSED';
    END IF;

END; //

DELIMITER ;

CREATE USER dev@'%' IDENTIFIED BY 'dev';
GRANT ALL PRIVILEGES ON secretform.* TO dev@'%';
