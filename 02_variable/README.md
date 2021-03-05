# **Variable Rules and regulations**
 - [x] **Variable declaration**:

- var foo int
- var foo int = 41
- foo:=42

- [x] **Can't redeclare variables,but can shadow them**.
- [x] **All variables must be used**.

- [x] **Vasibility**:
- lower case first letter for package scope
- upper case first letter to export
- no private scope
- [x] **Naming conversions**:
- Pascal or camel case *.
  -capitalize acronyms(HTTP,URL)
- As short as reasonable
  -longer names for longer lives

- [x] **Type Conversions**:
- destinationType(variable)
- use strconv package for strings
