#include <omp.h>
#include <stdio.h>

int main() {

	#pragma omp parallel
	{
		int ID = omp_get_thread_num();
		for(int i = 1+(100*ID); i<100*(ID+1); i++){
			printf("hello(%d)", ID);
			printf("world(%d)\n", ID);
		}
	}

    return 0;
}
