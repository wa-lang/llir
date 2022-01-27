@0 = global [12 x i8] c"hello llir!\00"

declare i32 @puts(i8* %0)

define i32 @main() {
0:
	%1 = getelementptr [12 x i8], [12 x i8]* @0, i32 0, i32 0
	%2 = call i32 @puts(i8* %1)
	ret i32 0
}

