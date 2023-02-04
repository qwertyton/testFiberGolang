# testFiberGolang
for exam

# mockup
```
create
    definer = root@localhost procedure mockup()
BEGIN
    DECLARE round INT default 0;
    DECLARE one INT;
    DECLARE two INT;
    test:
    LOOP
        select floor(0 + RAND() * 10) into one;
        select floor(0 + RAND() * 10) into two;
        insert into test values (one, two);
        set round = round + 1;
        IF round >= 1000 THEN
            commit;
            LEAVE test;
        END IF;
    END LOOP test;
END;
```