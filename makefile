git-add:
	git add .
	git commit -am '${cmt}'

test:
	echo ${cmt}

entermysql:
	docker exec -it mysql_fiber_gorm_rakamin mysql -u ADMIN -pSECRET rakamin_intern

runenv:
	docker compose up -d

run:
	docker compose up -d
	go run app/main.go

struct:
	gomodifytags -file ${file} -struct ${struct} -add-tags ${tags}

stop:
	docker compose stop

down:
	docker compose down -v

logs:
	docker compose logs -f