# MBSS Migration Tool

> Database migration tool for DMAA

## Prerequisites

* [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) CLI

## Quick Config

* Copy and setup dbparam
```
  cp dbparam.txt.example dbparam.txt
```

## Create new SQL file
```
  ./create.sh new_migration_file_name
```

## Migrate

* Up (migrate up to N files)
```
  ./migrate_up.sh N
```

* Down (migrate down to N files)
```
  ./migrate_down.sh N
```

* Goto (migrate all files (up/down) to specific version)
```
  ./migrate_goto.sh version
```

* Force (skip migrate specific version)
```
  ./migrate_force.sh version
```
