#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>
#include <string.h>

void comparision(char ch, va_list arguments);
int infof(const char *format, ...);
int warnf(const char *format, ...);
int errorf(const char *format, ...);
int panicf(const char *format, ...);

void textcolor(int color)
{	
	printf("\e[%dm", color);
}

int infof(const char *format, ...){
	textcolor(97);
	va_list arguments;
	va_start(arguments, format);

	while(*format != NULL){
		if(*format == '%'){
			format++;
			comparision(*format++, arguments);
		}else{
			printf("%c", *format);
			format++;
		}
	}

	va_end(arguments);
	textcolor(0);
	return 0;
}

int warnf(const char *format, ...){
	textcolor(93);
	va_list arguments;
	va_start(arguments, format);

	while(*format != NULL){
		if(*format == '%'){
			format++;
			comparision(*format++, arguments);
		}else{
			printf("%c", *format);
			format++;
		}
	}

	va_end(arguments);
	textcolor(0);
	return 0;
}

int errorf(const char *format, ...){
	textcolor(91);
	textcolor(1);
	va_list arguments;
	va_start(arguments, format);

	while(*format != NULL){
		if(*format == '%'){
			format++;
			comparision(*format++, arguments);
		}else{
			printf("%c", *format);
			format++;
		}
	}

	va_end(arguments);
	textcolor(0);
	return 0;
}

int panicf(const char *format, ...){
	textcolor(0);
	abort();
	return 0;
}

void comparision(char ch, va_list arguments){
	int d;
	char c, *s;
	switch (ch){
		case 's':              /* string */
            s = va_arg(arguments, char *);
            printf("%s", s);
            break;
        case 'd':              /* int */
            d = va_arg(arguments, int);
            printf("%d", d);
            break;
        case 'c':              /* char */
            /* need a cast here since va_arg only
               takes fully promoted types */
            c = (char) va_arg(arguments, int);
            printf("%c", c);
            break;
       }
}