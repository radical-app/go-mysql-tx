# Golang Mysql transactions helper

Working with golang and mysql in the good way.

- Transactions are awesome tool for sql developers.
- Context is an awesome tool for GO developers.

### Why not an orm?

- Too slow.
- Too complex.

### why not raw sql?

- Developer experience is poor.
- `sql` module does not enforce good practices.

### 0 configure vars 
    

    
### The flow

    // the main
    c := mysql.ConfigFromEnvs("TEST") 
    
    db, err := mysql.Open(c, ctx)
    if err != nil {
          // do smtg clever 
    }
    //db.SetMaxLifetimeMins(15)
  
    // usually in the request/response ctx 
    // ctx := context.Background()
    tx, err := mysql.TxCreate(db, ctx)
    if mysql.IsErrorRollback(err, tx) {
        // do smtg clever 
    }
   
    incremental, err := mysql.TxPush(tx, ctx, "insert into MYDB (name) values (?)", "namearg")
    if mysql.IsErrorRollback(err, tx) {
        // do smtg clever is already rollbacked
    }
    fmt.Print(incremental)
    // -----
    // multiple insert on single transaction
    // -----
    incremental, err := mysql.TxPush(tx, ctx, "insert into MYDB (name) values (?)", "namearg2")
    if mysql.IsErrorRollback(err, tx) {
        // do smtg clever is already rollbacked
    }
    fmt.Print(incremental)
    err = tx.Commit()
    if err != nil {
        // do smtg clever 
    }
  
  