import { ResponseViewModel } from '../models/responses/response-viewmodel';
import axios from 'axios';

interface AxiosErrorResponse {
  status: number;
}
interface AxiosError {
  response: AxiosErrorResponse;
}
export class HttpBaseService {
  public static init(
    errorCallback: (statusCode: number) => void,
    loadingCallback: (really: boolean) => void
  ) {
    console.log('endpoint', process.env.apiEndPoint);
    // axios.defaults.baseURL = 'http://localhost:8080';
    axios.defaults.headers = {
      'content-type': 'application/json'
    };
    axios.interceptors.request.use(
      (config: any) => {
        loadingCallback(true);
        return config;
      },
      (error: any) => {
        loadingCallback(false);
        return Promise.reject(error);
      }
    );
    axios.interceptors.response.use(
      (response: any) => {
        loadingCallback(false);
        return response;
      },
      (error: AxiosError) => {
        loadingCallback(false);
        const errorPromise = Promise.reject(error);
        if (error.response) {
          errorCallback(error.response.status);
        }
        return errorPromise;
      }
    );
  }

  public static async post<Req, Res>(url: string, body: Req) {
    try {
      const response = await axios.post<ResponseViewModel<Res>>(
        url,
        JSON.stringify(body)
      );
      return response.data;
    } catch (e) {
      return Promise.reject(e);
    }
  }
}
