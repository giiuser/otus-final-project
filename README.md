# otus-final-project

## Превьювер изображений

Поддерживаемые расширения: jpg, png, gif.

## Использование

```console
make run
```

## Запуск тестов

```console
make test
```

```console
make test-race
```

## Пример

После запуска в консоли, открыть URL в браузере

```console
http://localhost:8082/fill/300/200/www.audubon.org/sites/default/files/a1_1902_16_barred-owl_sandra_rothenberg_kk.jpg
```

## Переменные окружения (по умолчанию)

```console
PORT=8082
MAX_FILE_SIZE=5242880
CACHE_DIR=.cache
CACHE_SIZE=10
