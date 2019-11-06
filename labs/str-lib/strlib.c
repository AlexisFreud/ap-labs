#include <stdlib.h>
#include <stdio.h>

int mystrlen(char *str){
	int len = 0;
	while(*str != NULL){
		len++;
		str++;
	}
    return len;
}

char *mystradd(char *origin, char *addition){
	int newLenght = mystrlen(origin) + mystrlen(addition);
	char *newString = malloc(sizeof(char)*newLenght);
	for(int i = 0; i < mystrlen(origin); i++){
		newString[i] = origin[i];
	}
	for(int j = 0; j < mystrlen(addition); j++){
		newString[j+mystrlen(origin)] = addition[j];
	}
    return newString;
}

int mystrfind(char *origin, char *substr){
    for (int i = 0; i < mystrlen(origin); ++i)
    {
    	for(int j = 0; j < mystrlen(substr); j++){
    		if(origin[i+j] != substr[j]){
    			break;
    		}else if(j == mystrlen(substr)-1){
    			return 1;
    		}
    	}
    }
    return 0;
}