# Примитивы синхронизации

# sync/atomic

Реализацию примитивов из пакета `sync/atomic` можно найти в исходном коде Go:

`src/sync/atomic/doc.go` -> `src/sync/atomic/asm.s` -> `src/internal/runtime/atomic/atomic_amd64.go` ->
    `src/internal/runtime/atomic/atomic_amd64.s`

Описание префикса `LOCK` можно найти здесь: https://www.felixcloutier.com/x86/lock

## Гарантии atomic

Гарантии предоставляемые пакетом `atomic` описаны в документе The Go Memory Model: https://go.dev/ref/mem#

Полезно ознакомиться с этим документом: https://research.swtch.com/gomm.pdf . Это описание относительно недавних изменений
в описании модели памяти.

# sync.Mutex

См. код в `src/sync/mutex.go` (исходный код Go).

Вопрос: почему мьютексы нельзя копировать ?

## sync.RWMutex

См. код в `src/sync/rwmutex.go` (исходный код Go).

Вопрос: когда стоит использовать RWMutex ?

См. бенчмарки в `internal/mmutex/rwmutex_test.go`

# sync.Once

Попробуем написать `once` самостоятельно. Вы можете использовать атомарные операции и мьютексы.

Почему код в `internal/once/incorrect.go` не соответствует спецификации ?

См. код в `src/sync/mutex.go` (исходный код Go).

Вопрос: зачем исползовать одновременно `atomic.Uint32` и `Mutex` ?

# sync.Pool

См. код в `src/sync/pool.go` (исходный код Go).

Для лучшего понимания устройства планировщика в Go можно обратиться к https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html

# Каналы

См. `src/runtime/chan.go` (исходный код Go).

Особенно нас интересуют функции `chansend` и `chanrecv`.

Также стоит обратить внимание на `recvDirect` и `sendDirect`.

# Пакет context

Попробуем написать `WithCancel()` для контекста самостоятельно. См. файл `internal/context/cancel.go`. Вы можете использовать любые примитивы синхронизации.

Решение см. в `internal/context/solution.go`.
