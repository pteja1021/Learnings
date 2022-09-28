CREATE OR REPLACE FUNCTION CHECK_USERNAME()
RETURNS TRIGGER
AS $BODY$
BEGIN
IF LENGTH(NEW.NAME) < 4 THEN
    raise exception 'length cant be less than 4';
ELSE RETURN NEW;
END IF;
END;
$BODY$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER CHECK_USER_NAME 
BEFORE INSERT ON "public"."users"
FOR EACH ROW 
EXECUTE PROCEDURE CHECK_USERNAME();
