export interface Queue {
  id?: number;
  code: string;
  type: string;
  timeStamp: string;
  status?: string;
}

export interface PostQueue {
  type: string;
  name: string;
  tel: string;
}

export interface QueueDetail {
  code: string;
  qr?: string;
  barCode?: string;
  name?: string;
  tel?: string;
  timestamp: string;
}

export interface HttpResponse {
  Code: string;
  Date: string;
  Type: string;
  Name: string;
  Tel: string;
}

export interface DateAndTime {
  date: string;
  time: string;
}
