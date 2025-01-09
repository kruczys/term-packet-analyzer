#ifndef INCLUDE_CAPTURE_H
#define INCLUDE_CAPTURE_H

#include "common.h"

pcap_t* create_pcap_handle(char* device, char* filter);
int get_link_header_len(pcap_t *handle);
void stop_capture(int sig_number);

#endif // INCLUDE_CAPTURE_H
