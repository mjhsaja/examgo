# Linked List (Go)

Implementasi *singly linked list* generik dalam Go. Kode ada di [linked-list.go](linked-list.go).

## Konsep

Linked list adalah struktur data linear yang elemennya (disebut *node*) tidak disimpan berurutan di memori. Setiap node menyimpan nilai dan pointer ke node berikutnya.

```
head                               tail
 │                                  │
 ▼                                  ▼
┌────┬───┐   ┌────┬───┐   ┌────┬───┐
│ 10 │ ●─┼──▶│ 20 │ ●─┼──▶│ 30 │ ∅ │
└────┴───┘   └────┴───┘   └────┴───┘
```

Dibandingkan slice, linked list unggul untuk menambah/menghapus di awal list (O(1)), tetapi lebih lambat untuk akses berdasarkan index (O(n)) karena harus menelusuri node satu per satu.

## Cara Menjalankan

Butuh Go 1.18 ke atas (karena memakai generics).

```bash
go run linked-list.go
```

## Struktur Data

```go
type Node[T any] struct {
    Value T
    Next  *Node[T]
}

type LinkedList[T comparable] struct {
    head *Node[T] // node pertama
    tail *Node[T] // node terakhir, agar PushBack O(1)
    size int      // jumlah elemen, agar Len O(1)
}
```

- `T` harus `comparable` karena `Remove`, `IndexOf`, dan `Contains` membandingkan nilai dengan `==`.
- `head`, `tail`, dan `size` tidak diekspor sehingga hanya bisa diubah lewat method.

## API

### Membuat list

| Fungsi | Keterangan |
|---|---|
| `NewLinkedList[T]() *LinkedList[T]` | Membuat list kosong. |

### Menambah elemen

| Method | Keterangan | Kompleksitas |
|---|---|---|
| `PushFront(v T)` | Menambah `v` di awal list. | O(1) |
| `PushBack(v T)` | Menambah `v` di akhir list. | O(1) |
| `Insert(index int, v T) error` | Menyisipkan `v` di posisi `index` (0 ≤ index ≤ Len). | O(n) |

### Menghapus elemen

| Method | Keterangan | Kompleksitas |
|---|---|---|
| `PopFront() (T, bool)` | Menghapus dan mengembalikan elemen pertama; `false` jika list kosong. | O(1) |
| `RemoveAt(index int) (T, error)` | Menghapus elemen di posisi `index` dan mengembalikan nilainya. | O(n) |
| `Remove(v T) bool` | Menghapus kemunculan pertama `v`; `false` jika tidak ditemukan. | O(n) |

### Mengakses dan mencari

| Method | Keterangan | Kompleksitas |
|---|---|---|
| `Get(index int) (T, error)` | Mengambil nilai di posisi `index`. | O(n) |
| `IndexOf(v T) int` | Posisi kemunculan pertama `v`, atau `-1`. | O(n) |
| `Contains(v T) bool` | Apakah `v` ada di list. | O(n) |
| `Len() int` | Jumlah elemen. | O(1) |
| `IsEmpty() bool` | Apakah list kosong. | O(1) |

### Utilitas

| Method | Keterangan | Kompleksitas |
|---|---|---|
| `Reverse()` | Membalik urutan list secara in-place. | O(n) |
| `ToSlice() []T` | Mengubah list menjadi slice. | O(n) |
| `String() string` | Format teks, contoh `[1 -> 2 -> 3]`. Dipakai otomatis oleh `fmt.Println`. | O(n) |

### Error

| Error | Kapan terjadi |
|---|---|
| `ErrIndexOutOfRange` | `Insert`, `RemoveAt`, atau `Get` dipanggil dengan index di luar jangkauan. |

Cek dengan `errors.Is(err, ErrIndexOutOfRange)`.

## Contoh Penggunaan

```go
list := NewLinkedList[int]()

list.PushBack(10)
list.PushBack(20)
list.PushFront(5)          // [5 -> 10 -> 20]

_ = list.Insert(2, 15)     // [5 -> 10 -> 15 -> 20]

v, _ := list.Get(1)        // 10
list.Remove(15)            // [5 -> 10 -> 20]
list.Reverse()             // [20 -> 10 -> 5]

first, ok := list.PopFront() // 20, true

if _, err := list.Get(10); err != nil {
    fmt.Println(err)       // index di luar jangkauan
}
```

List juga bisa dipakai untuk tipe lain, misalnya `NewLinkedList[string]()`.

## Cara Kerja Operasi Penting

### PushBack

Karena list menyimpan pointer `tail`, node baru langsung disambungkan ke `tail.Next` tanpa menelusuri list. Jika list kosong, node baru menjadi `head` sekaligus `tail`.

### Remove / RemoveAt

Menelusuri list sambil menyimpan node sebelumnya (`prev`). Saat target ditemukan, `prev.Next` diarahkan ke `target.Next` sehingga target terlepas dari rantai. Jika target adalah `tail`, maka `tail` dipindah ke `prev`.

```
sebelum:  [5] ──▶ [15] ──▶ [20]
sesudah:  [5] ─────────────▶ [20]
```

### Reverse

Menelusuri list sekali sambil membalik arah pointer `Next` setiap node, memakai tiga variabel: `prev`, `cur`, dan `next`. Di akhir, `head` lama menjadi `tail` dan node terakhir menjadi `head`.

```
sebelum:  [1] ──▶ [2] ──▶ [3]
sesudah:  [1] ◀── [2] ◀── [3]
           tail             head
```

## Catatan

- Struktur ini **tidak thread-safe**. Jika dipakai dari banyak goroutine, lindungi dengan `sync.Mutex`.
- Go sudah menyediakan *doubly linked list* di package standar [`container/list`](https://pkg.go.dev/container/list), tetapi tanpa generics. Implementasi ini dibuat untuk tujuan belajar.
