#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	int a, b[5] = {0, 1, 2, 1, 2};

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d", &a) != EOF)
		printf("%d\n", a / 5 + b[a % 5]);
	return 0;
}
