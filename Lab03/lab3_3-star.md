s
# Лабораторная работа №3 (3*) "Настройка CI/CD"

## Выполнили: 
Бевз Тимофей K34201, Загайнова Кристина K34201, Блохина Анастасия K34201, Балашов Матвей K34201

## Цель работы:
Реализация настроек CI/CD для автоматического запуска и сохранения результата запуска образа Docker после пуша файлов в репозиторий, а также работа с секретами посредством Hashicorp Vault.

## Задачи:
* Настройка CI/CD для автоматизации работы с образами Docker
* Проверка работы настроек
* Работа с секретами Hashicorp Vault

## Ход работы

### Настройка CI/CD на Github

1.  Сначала в локальном репозитории был запущен Hashicorp Vault. 

* Запуск Hashicorp Vault

![img1](./img/lab3_vault.jpg)

Данные в локальном Vault хранятся в защифрованном виде.

* Данные в локальном Vault

![img2](./img/lab3_vault_check.jpg)

2.  Затем по руководству от Github были настроены self runners для автоматического запуска образов Docker.

* Руководство по настройке self runners

![img3](./img/lab3_self_runner.jpg)

3. Список успешно настроенных self runners представлен на рисунке ниже.

* Руководство по настройке self runners

![img4](./img/lab3_runners.jpg)

4. Затем в локальном репозитории были проведены попытки запуска настроенных self runners с помощью средств Github Actions.

* Первый запуск образа

![img5](./img/lab3_gh_act.jpg)

Среди причин неудач можно выделить следующие:

* Установка локальных переменных для VAULT
* Отсутствие докера на self run машине

5. После ряда неудач и фикса ошибок удалось достичь успешного запуска образа Docker.

* Успешный запуск образа

![img6](./img/lab3_success.jpg)


6.  Ниже представлен Access Token для Dockerhub по итогу работы.

* Access Token

![img7](./img/lab3_acc_token.jpg)


7.  Ниже представлен код файла [lab3_star_build.yml](https://github.com/T1vz/itmo_clouds/blob/main/.github/workflows/lab3_star_build.yml) для автоматической сборки и публикации образа Docker.

* Имя процесса

```
name: Build and Publish
```

* Ветка и директории, в которой нужно отслеживать изменения
```
on:
  push:
    branches: [ "main" ]
    paths:
      - "Lab03/**"
      - ".github/workflows/**"
```

* Шаги запускаемого процесса

```
jobs:

  build:
    runs-on: self-hosted
    steps:
      - name: Checkout
        uses: actions/checkout@v3

      - name: Import Secrets
        uses: hashicorp/vault-action@v2
        with:
          url: http://127.0.0.1:8200
          tlsSkipVerify: true
          token: ${{ secrets.VAULT_TOKEN }}
          secrets: |
            secret/data/docker * | DOCKER_

      - name: Log in to Docker Hub
        uses: docker/login-action@v3
        with:
          username: ${{ env.DOCKER_LOGIN }}
          password: ${{ env.DOCKER_TOKEN }}

      - name: Build and Push image to docker.io
        uses: docker/build-push-action@v2
        with:
          context: Lab03
          push: true
          tags: ${{ env.DOCKER_LOGIN }}/${{ env.DOCKER_NAME }}:latest
```

## Вывод:
В результате выполнения лабораторной работы были изучены основы работы с CI/CD на Github, были настроены self runners для автоматического запуска образов Docker, была проведена установка и работа со средствами Hashicorp Vault.
Планируемые результаты были достигнуты.
