#include <stdio.h>

int max(int a, int b) {
	return (a > b) ? a : b;
}

int nrow(int a) {
	int p = 0, c = 0;
	while (a > c)
		c += ++p+1;
	return p;
}

int main(int argc, char *argv[])
{
	FILE *fp;
	int n = 0, m[5050] = {}, p;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d", &m[n]) != EOF)
		n++;
	p = n - 1;
	while (p) {
		int i, nr = nrow(p);
		for (i = 0; i < nr; i++)
			m[p - i - nr - 1] += max(m[p - i], m[p - i - 1]);
		p -= nr + 1;
	}
	printf("%d\n", m[0]);
	return 0;
}
