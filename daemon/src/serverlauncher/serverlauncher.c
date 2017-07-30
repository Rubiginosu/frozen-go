#define _GNU_SOURCE
#include <sched.h>
#include <sys/types.h>
#include <errno.h>
#include <sys/resource.h>
#include <unistd.h> 
#include<stdlib.h>
#include <stdio.h>
#include <string.h>
int main( int argc,char *argv[]){ 
struct rlimit rlim_new; 
int m=atoi(argv[2]); 
int uid=atoi(argv[1]); 
char* home=argv[3];
char* sv=argv[4];
rlim_new.rlim_cur = rlim_new.rlim_max =80;
setrlimit(RLIMIT_NPROC, &rlim_new);
rlim_new.rlim_cur = rlim_new.rlim_max =m*1024*1024; 
setrlimit(RLIMIT_AS, &rlim_new);
chroot(home);
chdir(home);
printf("%d %d\n",unshare(CLONE_NEWUTS|CLONE_NEWIPC|CLONE_FS|CLONE_FILES),errno);
//printf("%d %d\n",mount("proc", "./proc", "proc", 0, NULL),errno);
chdir(sv);
sethostname("cpera",strlen("cpera"));
while(setgid(uid)!=0) sleep(1);
while(setuid(uid)!=0) sleep(1);
setresuid(uid,uid,uid);
//argv[0]="su";
//argv[1]="fgo";
//argv[2]="-f";
//argv[3]="-c"; 
execvp(argv[5],&argv[5]);
printf("error: %d\n",errno);
}
