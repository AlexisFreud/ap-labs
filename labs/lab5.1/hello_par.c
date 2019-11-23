#include "omp.h"
#include <stdio.h>

int main() {

	#pragma omp parallel
	{
		int ID = omp_get_thread_num();
		for(int i = 1+(100*ID); i<100*(ID+1); i++){
			printf("Im %d and going %d\n", ID, i);
		}
	}

    return 0;
}
