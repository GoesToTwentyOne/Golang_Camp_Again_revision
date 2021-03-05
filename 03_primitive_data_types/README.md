# **Primitive Data Types**
 - [x] **Boolean type**:
- Value are true or false.
- Not an alias for other types(e.g. int)
- Zero value is false.

- [x] **Numeric types :**
- **1.Integers**
  - Signed integers
    - int type has varying size but min 32  bits.
    - 8 bit(int 8) through 64 bit (int64)
  -Unsinged integers 
    - 8 bit (byte and uint8) through 32 bit(uint32).
- Zero value is 0.
- Can't mix types in same family ! (int16+int32= error).

**2.Floating point numbers**
- Follow IEEE-754 standard.
- Zero value is 0.
- 32  and 64 bit versions of it.
- Literal styles
  - Decimal (2.5)
  - Exponential(13e18 or 2E10)
  - Mixed (13.7e12).


**3.Complex numbers**
- Zero value is (0+0i)
- 64 and 128 bit versions
- Built - in functions 
  - complex - make complex number from two floats.
  - real - get real part as float
  -imag - get imaginary part as float


**Text types**
- strings
  - UTF-8
  - immutable
  - can be concatenated with plus(+) operator.
- Rune 
  - UTF-32
  - Alias for int32
  - Special methods normally requred to process
    -.e.g strungs.Read #ReadRune

