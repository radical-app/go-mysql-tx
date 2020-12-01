# Golang fast SQL toolbox
### Named parameters, Transactions with Context

- Faster *named parameters*, no reflection!
- Simpler consistency - *Transactions* always commit or rollback! 
- SQL with *Context*, you should catch problems from the outside.
- Improving Developer experience enforcing Good practices.

## Named parameters 

```go main.go
skill := "go"
FetchRows(db, ctx, "SELECT me FROM devs WHERE skills = :skill", Named{"skill": skill})
```
fixed: `https://github.com/go-sql-driver/mysql/issues/561`

## Single or Multiple Transaction 

### Open a connection and Ping it || reuse your existing *sql.DB

```go
// main.go

c := toolbox.ConfigFromEnvs("TEST") 

db, err := toolbox.Open(c, ctx)
if err != nil {
      // do smtg clever 
}
//db.SetMaxLifetimeMins(15) optional set it
```

### Create a context || reuse the existing (from the http-framework maybe)
 
```go repo.go
// in repo.go

// ctx = context.Background() or reuse the ctx from the request/response framework
tx, err := toolbox.TxCreate(db, ctx) 
if toolbox.IsErrorRollback(err, tx) {
    // do smtg clever, tx is already rollbacked
}

// start using the transaction
incrementalOId, err := toolbox.TxPush(tx, ctx, "insert into `order` (name) values (:name)", toolbox.Named{"name": "value to insert"})
if toolbox.IsErrorRollback(err, tx) {
    // do smtg clever is already rollbacked
}

incrementalIId, err := toolbox.TxPush(tx, ctx, "insert into `item` (order_id) values (:order_id)", toolbox.Named{"order_id": incrementalOId})
if toolbox.IsErrorRollback(err, tx) {
    // do smtg clever is already rollbacked
}

if tx.Commit() != nil {
    // do smtg clever
}
```

 
  