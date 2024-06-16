# Build Your Own Google Keep

Problem statement https://codingchallenges.fyi/challenges/challenge-keep/
Writeup https://srinjoysantra.live/blog/go-keep

## Concepts

- [x] Directory structure
- [x] Dependency injection
- [x] Request interceptors
-  API
   -  [x] HTTP
      -  [x] Ouath2 with Auht0
      -  [x] Session based on cookie
- [x] Vendoring
- [x] Commenting
- [ ] Containerizing
- [ ] Logging
- [x] Error
- [ ] Tracing
- [ ] Documenting
- [ ] Unit Testing
- [ ] Integration testing

The general philosophy of it was not to use various HTTP rest client libraries like gin / echo / etc and use fully embrace `net/http`

## Resources

- Auth https://www.authgear.com/post/authentication-as-a-service
- Social login https://developer.auth0.com/resources/code-samples/api/standard-library/basic-authorization
- API best practices https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/
- https://12factor.net
- status code https://www.restapitutorial.com/httpstatuscodes.html
- rest best practices https://stackoverflow.blog/2020/03/02/best-practices-for-rest-api-design/#h-handle-errors-gracefully-and-return-standard-error-codes
- https://restfulapi.net/http-methods

### DB queries
https://www.postgresql.org/docs/16/index.html
```sql
CREATE TABLE Note(
    id UUID NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);
```

https://aviyadav231.medium.com/automatically-updating-a-timestamp-column-in-postgresql-using-triggers-98766e3b47a0

```sql
CREATE FUNCTION update_updated_at_note()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';
```

```sql
CREATE TRIGGER update_note_updated_at
    BEFORE UPDATE
    ON
        note
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_note();
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

### Auth
- [OAuth 2.0 and OpenID Connect (in plain English) YT video](https://www.youtube.com/watch?v=996OiexHze0)
- Auth0 alternatives
  - https://www.reddit.com/r/node/comments/17s0c6c/is_there_a_auth0_alternative_that_is_cost/

### Session handling
#### Session libraries
  - Gorilla session (Used in auth0 example)
  - https://github.com/alexedwards/scs
  - https://www.reddit.com/r/golang/comments/111i0ps/what_are_the_best_alternatives_to_gorilla_session/

#### Create your own session handling
  - https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/06.1.html
  - https://www.reddit.com/r/golang/comments/13yrkr6/authentication_with_nethttp_package/
  - https://github.com/Acebond/session
  - https://www.alexedwards.net/blog/working-with-cookies-in-go

### Deployment
 Where to deploy
  - https://www.reddit.com/r/golang/comments/rwjpw9/where_do_you_guys_deploy_your_hobby_projects/
