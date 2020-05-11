# Static Files

Example to create an atreugo static file server with basic middleware configuration.<br/>
If you want to see subdirecories and their files served, execute commands below.

```bash
mkdir -p nested/one
echo 'Nested Directory' > ./nested/dir.txt
echo 'One More' > ./nested/one/more.txt
```

### Routes:
- `/`
- `/main`
- `/readme`
- `/gitignore`

- `/static/default`
- `/static/middlewares`
- `/static/custom`
- `/static/readme`
- `/static/gitignore`
