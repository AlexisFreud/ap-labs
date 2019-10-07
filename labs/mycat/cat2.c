#include <stdio.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>

void sysFilecopy(FILE *ifp, FILE *ofp)
{
	int c;

    while ((c = getc(ifp)) != EOF)
        putc(c, ofp);
}

/* filecopy:  copy file ifp to file ofp */
void filecopy(int ifp, FILE *ofp)
{
	int c, n;
    char buffer[5000000];
    n = read(ifp, buffer, 5000000);
    int i = 0;
    while (buffer[i] != NULL){
        putc(buffer[i], ofp);
       	i++;
    }
}

/* cat:  concatenate files, version 2 */
int main(int argc, char *argv[])
{
    int fp; // File to be readed
    char *prog = argv[0];   /* program name for errors */
    if (argc == 1) { /* no args; copy standard input */
        sysFilecopy(stdin, stdout);
    }
    else 
	{
        while (--argc > 0)
        {
            if ((fp = open(*++argv, O_RDONLY, 0)) == -1) {
                fprintf(stderr, "%s: can′t open %s\n", prog, *argv);
                return 1;
            } 
            else 
            {
            	filecopy(fp, stdout);
                close(fp);
            }
        }
    }
    if (ferror(stdout)) {
        fprintf(stderr, "%s: error writing stdout\n", prog);
        return 2;
    }
    printf("\n");

    return 0;
}
