# Build Your Own Google Keep


https://www.authgear.com/post/authentication-as-a-service

```sql
CREATE TABLE Note(
    api_key UUID NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (api_key)
);
```

https://aviyadav231.medium.com/automatically-updating-a-timestamp-column-in-postgresql-using-triggers-98766e3b47a0
```sql
CREATE  FUNCTION update_updated_on_note()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_on = now();
    RETURN NEW;
END;
$$ language 'plpgsql';
```

```sql
CREATE TRIGGER update_note_updated_on
    BEFORE UPDATE
    ON
        note
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_on_note();
```

```sql
INSERT INTO Note VALUES('018de613-34ef-79ed-874b-2196c5d5167a', 'first note', '# Heading 1 ## Heading 2 body');
```

```sql
 SELECT * FROM note WHERE apiid='018de613-34ef-79ed-874b-2196c5d5167a';
```

```sql
UPDATE Note SET title = "Updated note", content="Why do I want to update?" WHERE apiid='018de613-34ef-79ed-874b-2196c5d5167a';
```