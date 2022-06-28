export interface Queue {
  id?: number;
  code: string;
  type: string;
  timeStamp: string;
  status?: string;
}

export interface PostQueue {
  type: string;
}

export interface QueueDetail {
  code: string;
  qr?: string;
  barCode?: string;
  timestamp: string;
}
