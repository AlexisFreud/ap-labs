#include <omp.h>

void matmultiplication();

int main() {
	matmultiplication();
}

void matmultiplication(){
	int Ndim, Pdim, Mdim;   /* A[N][P], B[P][M], C[N][M] */
    int i,j,k;
    double *A, *B, *C, cval, tmp, err, errsq;
	#pragma omp parallel for private(tmp, i, j, k)
		for (i=0; i<Ndim; i++){
			for (j=0; j<Mdim; j++){
				tmp = 0.0;
				for(k=0;k<Pdim;k++){
					/* C(i,j) = sum(over k) A(i,k) * B(k,j) */
					tmp += *(A+(i*Ndim+k)) * *(B+(k*Pdim+j));
				}
			*(C+(i*Ndim+j)) = tmp;
			}
		}
}
