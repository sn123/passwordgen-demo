import { HttpBaseService } from './http-base-service';
import { ResponseViewModel } from '@/models/responses/response-viewmodel';
import { PasswordResponse } from '@/models/responses/passwords-response';
import { GeneratorRequest } from '@/models/requests/generator-request';
export class GeneratorService {
  public static async generate(
    request: GeneratorRequest
  ): Promise<ResponseViewModel<PasswordResponse>> {
    return HttpBaseService.post<GeneratorRequest, PasswordResponse>(
      '/api/generate',
      request
    );
  }
}
