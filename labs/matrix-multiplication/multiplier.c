#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <pthread.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#include <sys/types.h>
#include "logger.h"

#define fileA "matA.dat"
#define fileB "matB.dat"
#define resultFile "result.dat"

int NUM_BUFFERS;
long **buffers;
pthread_mutex_t *mutexes;
long *result;
pthread_t threads[2000];

struct data{
	long *matA;
  long *matB;
  int row;
  int col;
  long *result;
};

int pthread_join(pthread_t thread, void **retval);
long *readMatrix(char *filename);
long *getColumn(int col, long *matrix); 
long *getRow(int row, long *matrix);
int getLock();
int releaseLock(int lock);
long dotProduct(long *vec1, long *vec2);
long *multiply(long *matA, long *matB);
int saveResultMatrix(long *result);
void *auxFunc(void *arg);

int main(int argc, char *argv[]){
	if(argc != 3)
		error("Should be 3 arguments.");
	NUM_BUFFERS = atoi(argv[2]);
	buffers = (long **)malloc(sizeof(long *)*NUM_BUFFERS);
	mutexes = (pthread_mutex_t *) malloc(sizeof(pthread_mutex_t)*NUM_BUFFERS);
	for (int i = 0; i < NUM_BUFFERS; i++) {
    	pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;
   		mutexes[i] = mutex;
    	pthread_mutex_init(&mutexes[i], NULL);
  	}


	long *matrixA = malloc(sizeof(long)*(NUM_BUFFERS/2));
	long *matrixB= malloc(sizeof(long)*(NUM_BUFFERS/2));
	matrixA = readMatrix(fileA);
	matrixB = readMatrix(fileB);
	multiply(matrixA, matrixB);
	saveResultMatrix(result);
	printf("End :)\n");
    return 0;
}

void cleanLine(char *line){
	int i = 0;
	while(line[i] != '\0'){
		line[i] = '\0';
		i++;
	}
}

long *readMatrix(char *filename){
	long *matrix = malloc(sizeof(long)*4000000);
	int file;
    char *buf = malloc(sizeof(char)*16000000);
    if ((file = open(filename, O_RDONLY, 0)) == -1)
    	error("can't open %s", filename);

    read(file, buf, 16000000);
    int i = 0;
    int startLine = 0;
    char *line = malloc(sizeof(char)*10);
    int current = 0;
    while(buf[i] != '\0'){
   		if(buf[i] == '\n'){
   			matrix[current] = strtol(line, NULL, 10);
   			current++;
   			startLine = 0;
   			cleanLine(line);
   		}else{
   			line[startLine] = buf[i];
   			startLine++;
   		}
    	i++;
    }
	return matrix;
}

long *getColumn(int col, long *matrix){
	long *auxMatrix = malloc(sizeof(long)*2000);
	for(int i = 0; i < 2000; i++){
		auxMatrix[i] = matrix[(i*2000)+col];
	}
	return auxMatrix;
} 

long * getRow(int row, long *matrix){
	int startRow = row*2000;
	long *auxMatrix = malloc(sizeof(long)*2000);
	for(int i = 0; i < 2000; i++){
		auxMatrix[i] = matrix[i+startRow];
	}
	return auxMatrix;
}

int getLock(){
	for (int i = 0; i < NUM_BUFFERS; i++) {
    	if (pthread_mutex_lock(&mutexes[i]) == 0)
    		return i;
	}
	return -1;
}

int releaseLock(int lock){
	if (pthread_mutex_unlock(&mutexes[lock]) == 0) {
    	return 0;
  	}
  return -1;
}

long dotProduct(long *vec1, long *vec2){
	long result;
	for(int i = 0; i < 2000; i++){
		result += vec1[i]*vec2[i];
	}
	return result;
}

long * multiply(long *matA, long *matB){
	for (size_t i = 0; i < 2000; i++) {
    	for (size_t j = 0; j < 2000; j++) {
      		struct data *datos;
      		datos =	malloc(sizeof(struct data));
      		datos->matA = matA;
      		datos->matB = matB;
      		datos->row = i + 1;
      		datos->col = j + 1;
      		datos->result = result;
      		pthread_create(&threads[j], NULL, auxFunc, (void *)datos);
    	}

    	for (size_t j = 0; j < 2000; j++)
      		pthread_join(threads[j], NULL);
    	fflush(stdout);
    }
	return 0;
}

int saveResultMatrix(long *result){
	int file;
    if ((file = open(resultFile, O_WRONLY, 0)) == -1)
    	error("can't open %s", resultFile);

    for(int i = 0; i < 4000000; i++){
    	char* line = malloc(sizeof(char)*10);
    	sprintf(line, "%ld", result[i]);
    	write(file, line, strlen(line));
    	write(file, "\n", 1);
    }
	return 0;
}

void *auxFunc(void *arg) {
  struct data *datos = (struct data *)arg;
  long index;
  printf("creado\n");

  int lock1, lock2;
  while ((lock1 = getLock()) == -1);
  while ((lock2 = getLock()) == -1);
  buffers[lock1] = getRow(datos->row, datos->matA);
  buffers[lock2] = getColumn(datos->col, datos->matB);

  index = ((((datos->row - 1) * 2000) + datos->col) - 1);
  datos->result[index] = dotProduct(buffers[lock1], buffers[lock2]);

  while (releaseLock(lock1) != 0);
  while (releaseLock(lock2) != 0);
  return NULL;
}