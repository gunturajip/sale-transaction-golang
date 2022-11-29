# Tugas Akhir Rakamin Evermos Virtual Internship

### How to config MySQL DSN   
https://levelup.gitconnected.com/build-a-rest-api-using-go-mysql-gorm-and-mux-a02e9a2865ee

### API Daerah
https://www.emsifa.com/api-wilayah-indonesia/

### How to use makefile command
```bash
make {{command}}

# ex:
make run
```

### Check Relation
```sql
select * from INFORMATION_SCHEMA.TABLE_CONSTRAINTS;


SELECT 
  TABLE_NAME,COLUMN_NAME,CONSTRAINT_NAME, REFERENCED_TABLE_NAME,REFERENCED_COLUMN_NAME
FROM
  INFORMATION_SCHEMA.KEY_COLUMN_USAGE
WHERE
  REFERENCED_TABLE_SCHEMA = 'rakamin_intern' 
```